using System.Text.Json;

using Microsoft.AspNetCore.Diagnostics;
using Microsoft.AspNetCore.HttpLogging;

namespace Api.Extensions;

public static class MiddlewareExtensions
{
    private static readonly PathString HealthcheckEndpoint = "/health";

    public static void AddMiddleware(this IServiceCollection services, bool isDevelopment)
    {
        services.AddHealthChecks();
        services.AddValidation();
        // Every error response is an RFC 9457 Problem Details object, matching the Go starter.
        services.AddProblemDetails(options =>
            options.CustomizeProblemDetails = context =>
            {
                // Replace the framework's default RFC 9110 links; keep any type an endpoint sets itself.
                if (context.ProblemDetails.Type is null ||
                    context.ProblemDetails.Type.StartsWith("https://tools.ietf.org/html/rfc9110", StringComparison.Ordinal))
                {
                    context.ProblemDetails.Type = "about:blank";
                }

                // Report validation errors under the JSON (camelCase) field names clients send.
                if (context.ProblemDetails is HttpValidationProblemDetails validation)
                {
                    validation.Errors = validation.Errors.ToDictionary(
                        error => JsonNamingPolicy.CamelCase.ConvertName(error.Key),
                        error => error.Value);
                }
            });
        services.AddCors(options =>
        {
            options.AddPolicy("AllowAll", policy =>
            {
                policy.AllowAnyOrigin()
                    .AllowAnyMethod()
                    .AllowAnyHeader();
            });
        });
        services.AddRateLimiting();
        services.AddHttpLogging(o =>
        {
            o.CombineLogs = true;
            o.LoggingFields = isDevelopment
                ? HttpLoggingFields.ResponseBody | HttpLoggingFields.ResponseHeaders
                : HttpLoggingFields.RequestProperties | HttpLoggingFields.ResponsePropertiesAndHeaders;
        });
    }

    public static void UseMiddleware(
        this WebApplication app)
    {
        app.UseExceptionHandler(new ExceptionHandlerOptions
        {
            // Malformed requests throw BadHttpRequestException in Development; keep their 4xx status.
            StatusCodeSelector = exception => exception is BadHttpRequestException badRequest
                ? badRequest.StatusCode
                : StatusCodes.Status500InternalServerError
        });
        app.UseStatusCodePages();
        app.UseHealthChecks(HealthcheckEndpoint);
        app.UseHttpLogging();
        app.UseCors("AllowAll");
        app.UseRateLimiter();
    }
}

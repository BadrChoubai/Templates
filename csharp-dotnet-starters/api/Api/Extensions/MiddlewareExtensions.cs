using Microsoft.AspNetCore.HttpLogging;

namespace Api.Extensions;

public static class MiddlewareExtensions
{
    private static readonly PathString HealthcheckEndpoint = "/health";

    public static void AddMiddleware(this IServiceCollection services, bool isDevelopment)
    {
        services.AddHealthChecks();
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
        app.UseHealthChecks(HealthcheckEndpoint);
        app.UseHttpLogging();
        app.UseCors("AllowAll");
        app.UseRateLimiter();
    }
}

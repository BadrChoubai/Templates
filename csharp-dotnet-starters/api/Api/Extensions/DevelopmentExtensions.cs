namespace Api.Extensions;

public static class DevelopmentExtensions
{
    public static void AddDocumentation(this IServiceCollection services)
    {
        services.AddOpenApi();
    }


    public static void UseDocumentation(this WebApplication app)
    {
        app.MapOpenApi();
        app.UseSwaggerUI(options =>
        {
            options.SwaggerEndpoint("/openapi/v1.json", "v1");
            options.RoutePrefix = string.Empty;
        });
    }
}

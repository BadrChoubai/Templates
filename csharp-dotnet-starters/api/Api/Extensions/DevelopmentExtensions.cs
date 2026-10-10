namespace Api.Extensions;

public static class DevelopmentExtensions
{
    public static void AddDocumentation(this IServiceCollection services)
    {
        services.AddSwaggerGen();
        services.AddOpenApi();
    }


    public static void UseDocumentation(this WebApplication app)
    {
        app.MapOpenApi("/spec.json");
        app.UseSwagger();
        app.UseSwaggerUI(options =>
        {
            options.SwaggerEndpoint("/swagger/v1/swagger.json", "v1");
            options.RoutePrefix = string.Empty;
        });
    }
}
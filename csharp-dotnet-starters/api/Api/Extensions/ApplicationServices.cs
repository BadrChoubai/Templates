using Asp.Versioning;

using Api.Repository;
using Api.Service;

namespace Api.Extensions;

/// <summary>
/// class <c>AddApplicationServices</c> configures core application dependencies
/// </summary>
public static class ApplicationServices
{
    public static void AddServices(
        this IServiceCollection services
    )
    {
        services.AddEndpointsApiExplorer();
        services.AddRouting(options => options.LowercaseUrls = true);
        services.AddApiVersioning(options =>
        {
            options.DefaultApiVersion = new ApiVersion(1, 0);
            options.AssumeDefaultVersionWhenUnspecified = true;
            options.ReportApiVersions = true;
        });

        // Register Services
        services.AddScoped<IDataRepository, DataRepository>();
        services.AddScoped<IApiService, ItemsService>();
        services.AddTransient<DataSeeder>();
    }

    public static void AddDatabase(
        this IServiceCollection services,
        IConfiguration configuration)
    {
        var connectionString = configuration.GetConnectionString("DbConnectionString") ?? "Data Source=.db/api.db";
        services.AddSqlite<ApiDbContext>(connectionString);
    }
}
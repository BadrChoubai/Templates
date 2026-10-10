using Api.Repository;
using Api.Service;

namespace Api.Extensions;

/// <summary>
/// class <c>AddApplicationServices</c> configures core application dependencies
/// </summary>
public static class ApplicationServices
{
    private const string DefaultConnectionString = "Data Source=.db/api.db";

    public static void AddServices(
        this IServiceCollection services
    )
    {
        services.AddRouting(options => options.LowercaseUrls = true);

        // Register Services
        services.AddScoped<IDataRepository, DataRepository>();
        services.AddScoped<IApiService, ItemsService>();
        services.AddTransient<DataSeeder>();
    }

    public static void AddDatabase(
        this IServiceCollection services,
        IConfiguration configuration)
    {
        var connectionString = configuration.GetConnectionString("DbConnectionString") ?? DefaultConnectionString;
        services.AddSqlite<ApiDbContext>(connectionString);
    }
}

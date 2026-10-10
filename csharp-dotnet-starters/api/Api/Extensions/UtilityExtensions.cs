using Api.Repository;

using Microsoft.Data.Sqlite;
using Microsoft.EntityFrameworkCore;

namespace Api.Extensions
{
    public static class UtilityExtensions
    {
        /// <summary>
        /// Applies pending EF Core migrations, creating the SQLite database directory if needed.
        /// </summary>
        public static void MigrateDatabase(this IHost host)
        {
            using var scope = host.Services.CreateScope();
            var context = scope.ServiceProvider.GetRequiredService<ApiDbContext>();

            var dataSource = new SqliteConnectionStringBuilder(context.Database.GetConnectionString()).DataSource;
            var directory = Path.GetDirectoryName(dataSource);
            if (!string.IsNullOrEmpty(directory))
            {
                Directory.CreateDirectory(directory);
            }

            context.Database.Migrate();
        }

        public static void SeedData(this IHost host)
        {
            using var scope = host.Services.CreateScope();
            var service = scope.ServiceProvider.GetRequiredService<DataSeeder>();
            service.Seed();
        }
    }
}

namespace Api.Extensions
{
    public static class UtilityExtensions
    {
        public static void SeedData(this IHost host)
        {
            using var scope = host.Services.CreateScope();
            var service = scope.ServiceProvider.GetRequiredService<DataSeeder>();
            service.Seed();
        }
    }
}
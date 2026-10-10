using Api.Model;

using Microsoft.EntityFrameworkCore;

namespace Api.Repository
{
    public class ApiDbContext : DbContext
    {
        public ApiDbContext()
        { }

        public ApiDbContext(DbContextOptions<ApiDbContext> options) : base(options)
        { }

        public DbSet<Item> Items => Set<Item>();
    }
}

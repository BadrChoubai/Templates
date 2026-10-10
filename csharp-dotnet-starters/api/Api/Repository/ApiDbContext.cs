using Microsoft.EntityFrameworkCore;

using Api.Model;

namespace Api.Repository
{
    public class ApiDbContext : DbContext
    {
        public ApiDbContext()
        { }

        public ApiDbContext(DbContextOptions<ApiDbContext> options) : base(options)
        { }

        public DbSet<Item> Items { get; set; }
    }
}

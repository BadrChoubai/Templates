using Api.Model;

using Microsoft.EntityFrameworkCore;

namespace Api.Repository
{
    public class DataRepository(ApiDbContext dbContext) : IDataRepository
    {
        public async Task<List<Item>> GetItems()
        {
            return await dbContext.Items.ToListAsync();
        }

        public async Task<Item?> GetItem(Guid id)
        {
            return await dbContext.Items.FirstOrDefaultAsync(i => i.Id == id);
        }

        public async Task AddItem(Item item)
        {
            dbContext.Items.Add(item);
            await dbContext.SaveChangesAsync();
        }
    }
}

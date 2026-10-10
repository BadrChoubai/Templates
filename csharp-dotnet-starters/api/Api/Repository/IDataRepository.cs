using Api.Model;

namespace Api.Repository;

public interface IDataRepository
{
    Task<List<Item>> GetItems();
    Task<Item?> GetItem(Guid itemId);
    Task AddItem(Item item);
}

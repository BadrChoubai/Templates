using Api.Model;

namespace Api.Service;

public interface IApiService
{
    Task<List<Item>> GetAll();
    Task<Item?> GetItem(Guid id);
    Task<Item> CreateItem(CreateItemRequest request);
}

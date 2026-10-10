using Api.Model;
using Api.Repository;

namespace Api.Service;

public class ItemsService : IApiService
{
    private readonly IDataRepository _repository;
    private readonly ILogger<ItemsService> _logger;

    public ItemsService(IDataRepository repository, ILogger<ItemsService> logger)
    {
        _repository = repository;
        _logger = logger;
    }

    public async Task<List<Item>> GetAll()
    {
        _logger.LogInformation("Getting all items");
        return await _repository.GetItems();
    }

    public async Task<Item?> GetItem(Guid id)
    {
        _logger.LogInformation("Getting an item");
        return await _repository.GetItem(id);
    }

    public async Task<Item> CreateItem(CreateItemRequest request)
    {
        var item = new Item
        {
            Name = request.Name ?? string.Empty,
            Description = request.Description ?? string.Empty
        };

        _logger.LogInformation("Creating an item");
        await _repository.AddItem(item);
        return item;
    }
}

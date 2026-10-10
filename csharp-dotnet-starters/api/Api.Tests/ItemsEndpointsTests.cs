using System.Net;
using System.Net.Http.Json;

using Api.Model;

using Microsoft.Extensions.DependencyInjection;

namespace Api.Tests;

public class ItemsEndpointsTests : IClassFixture<ApiFactory>
{
    private static readonly Guid SeededId = new("c04ab9a0-e398-471d-8cfe-95317ccf0a20");

    private readonly HttpClient _client;

    public ItemsEndpointsTests(ApiFactory factory)
    {
        _client = factory.CreateClient();

        // The app applies migrations on startup; seeding is idempotent.
        using var scope = factory.Services.CreateScope();
        scope.ServiceProvider.GetRequiredService<DataSeeder>().Seed();
    }

    [Fact]
    public async Task Health_ReturnsHealthy()
    {
        var response = await _client.GetAsync("/health", TestContext.Current.CancellationToken);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Equal("Healthy", await response.Content.ReadAsStringAsync(TestContext.Current.CancellationToken));
    }

    [Fact]
    public async Task GetItems_ReturnsSeededItems()
    {
        var items = await _client.GetFromJsonAsync<List<Item>>("/items", TestContext.Current.CancellationToken);

        Assert.NotNull(items);
        Assert.Equal(2, items.Count);
        Assert.Contains(items, item => item.Id == SeededId);
    }

    [Fact]
    public async Task GetItem_ReturnsItem()
    {
        var item = await _client.GetFromJsonAsync<Item>($"/items/{SeededId}", TestContext.Current.CancellationToken);

        Assert.NotNull(item);
        Assert.Equal("First item", item.Name);
        Assert.Equal("An example item", item.Description);
    }

    [Theory]
    [InlineData("00000000-0000-0000-0000-000000000000")]
    [InlineData("not-a-guid")]
    public async Task GetItem_ReturnsNotFound_ForUnknownOrInvalidId(string id)
    {
        var response = await _client.GetAsync($"/items/{id}", TestContext.Current.CancellationToken);

        Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
    }
}

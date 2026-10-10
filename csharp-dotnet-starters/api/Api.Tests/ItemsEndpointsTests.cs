using System.Net;
using System.Net.Http.Json;
using System.Text;
using System.Text.Json;

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

    private static CancellationToken CancellationToken => TestContext.Current.CancellationToken;

    [Fact]
    public async Task Health_ReturnsHealthy()
    {
        var response = await _client.GetAsync("/health", CancellationToken);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Equal("Healthy", await response.Content.ReadAsStringAsync(CancellationToken));
    }

    [Fact]
    public async Task GetItems_ReturnsSeededItems()
    {
        var items = await _client.GetFromJsonAsync<List<Item>>("/items", CancellationToken);

        // The database is shared by the tests in this class, so other tests may add items.
        Assert.NotNull(items);
        Assert.Contains(items, item => item.Id == SeededId);
    }

    [Fact]
    public async Task GetItem_ReturnsItem()
    {
        var item = await _client.GetFromJsonAsync<Item>($"/items/{SeededId}", CancellationToken);

        Assert.NotNull(item);
        Assert.Equal("First item", item.Name);
        Assert.Equal("An example item", item.Description);
    }

    [Theory]
    [InlineData("00000000-0000-0000-0000-000000000000")]
    [InlineData("not-a-guid")]
    public async Task GetItem_ReturnsNotFoundProblem_ForUnknownOrInvalidId(string id)
    {
        var response = await _client.GetAsync($"/items/{id}", CancellationToken);

        await AssertProblem(response, HttpStatusCode.NotFound);
    }

    [Fact]
    public async Task CreateItem_ReturnsCreatedItem()
    {
        var response = await _client.PostAsJsonAsync(
            "/items",
            new { name = "New item", description = "Created in a test" },
            CancellationToken);

        Assert.Equal(HttpStatusCode.Created, response.StatusCode);
        var item = await response.Content.ReadFromJsonAsync<Item>(CancellationToken);
        Assert.NotNull(item);
        Assert.Equal("New item", item.Name);
        Assert.Equal($"/items/{item.Id}", response.Headers.Location?.OriginalString);

        var stored = await _client.GetFromJsonAsync<Item>($"/items/{item.Id}", CancellationToken);
        Assert.Equal("Created in a test", stored?.Description);
    }

    public static TheoryData<string, string[]> InvalidCreateRequests => new()
    {
        { """{"description": "No name"}""", ["name"] },
        { """{"name": "   "}""", ["name"] },
        { $$"""{"name": "{{new string('a', 256)}}"}""", ["name"] },
        { $$"""{"name": "", "description": "{{new string('a', 1025)}}"}""", ["name", "description"] },
    };

    [Theory]
    [MemberData(nameof(InvalidCreateRequests))]
    public async Task CreateItem_ReturnsValidationProblem_ForInvalidRequest(string body, string[] fields)
    {
        var response = await _client.PostAsync(
            "/items",
            new StringContent(body, Encoding.UTF8, "application/json"),
            CancellationToken);

        var problem = await AssertProblem(response, HttpStatusCode.BadRequest);
        var errorFields = problem.GetProperty("errors").EnumerateObject().Select(error => error.Name);
        Assert.Equal(fields.Order(), errorFields.Order());
    }

    [Fact]
    public async Task CreateItem_ReturnsBadRequestProblem_ForMalformedJson()
    {
        var response = await _client.PostAsync(
            "/items",
            new StringContent("""{"name":""", Encoding.UTF8, "application/json"),
            CancellationToken);

        await AssertProblem(response, HttpStatusCode.BadRequest);
    }

    /// <summary>
    /// Asserts the response is an RFC 9457 Problem Details object with the given status.
    /// </summary>
    private static async Task<JsonElement> AssertProblem(HttpResponseMessage response, HttpStatusCode status)
    {
        Assert.Equal(status, response.StatusCode);
        Assert.Equal("application/problem+json", response.Content.Headers.ContentType?.MediaType);

        var problem = await response.Content.ReadFromJsonAsync<JsonElement>(CancellationToken);
        Assert.Equal("about:blank", problem.GetProperty("type").GetString());
        Assert.Equal((int)status, problem.GetProperty("status").GetInt32());
        return problem;
    }
}

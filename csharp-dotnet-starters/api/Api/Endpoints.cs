using Api.Extensions;
using Api.Model;
using Api.Service;

namespace Api;

internal static class Endpoints
{
    public static void MapItemsApi(
        this IEndpointRouteBuilder routes)
    {
        var group = routes.MapGroup("/items");
        group.RequirePerUserRateLimit();

        group.WithTags("Items");

        group.MapGet("/", async (IApiService service) => await service.GetAll());
        group.MapGet("/{id:guid}", async (IApiService service, Guid id) =>
            await service.GetItem(id) is { } item
                ? Results.Ok(item)
                : Results.Problem(statusCode: StatusCodes.Status404NotFound, detail: "item not found"));
        group.MapPost("/", async (IApiService service, CreateItemRequest request) =>
        {
            var item = await service.CreateItem(request);
            return TypedResults.Created($"/items/{item.Id}", item);
        });
    }
}

using Api;
using Api.Extensions;

var builder = WebApplication.CreateBuilder(args);
var config = builder.Configuration;

builder.Services.AddServices();
builder.Services.AddDatabase(config);
builder.Services.AddMiddleware(
    builder.Environment.IsDevelopment()
);

// Configure Documentation
if (builder.Environment.IsDevelopment())
{
    builder.Services.AddDocumentation();
}

var app = builder.Build();
app.UseMiddleware();

if (app.Environment.IsDevelopment())
{
    app.UseDocumentation();
}

// Setup application endpoints
Endpoints.MapItemsApi(app);

app.MigrateDatabase();

if (args.Length == 1 && args[0].Equals("seed", StringComparison.OrdinalIgnoreCase))
{
    app.SeedData();
    return;
}

app.Run();

// Exposes the implicit Program class to WebApplicationFactory in Api.Tests.
public partial class Program;

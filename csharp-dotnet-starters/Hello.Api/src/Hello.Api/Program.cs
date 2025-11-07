using System.Diagnostics;
using System.Diagnostics.Metrics;
using Microsoft.Extensions.Diagnostics.HealthChecks;
using OpenTelemetry.Metrics;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;
using Serilog;

var builder = WebApplication.CreateBuilder(args);

// Load appsettings.json + environment-specific overrides
builder.Configuration
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .AddJsonFile($"appsettings.{builder.Environment.EnvironmentName}.json", optional: true, reloadOnChange: true)
    .AddEnvironmentVariables();

// Serilog setup
Log.Logger = new LoggerConfiguration()
    .ReadFrom.Configuration(builder.Configuration)
    .Enrich.FromLogContext()
    .WriteTo.Console(outputTemplate: "{Timestamp:yyyy-MM-dd HH:mm:ss} [{Level:u3}] {Message:lj}{NewLine}{Exception}")
    .CreateLogger();

builder.Host.UseSerilog();

// Add services
builder.Services.AddHealthChecks()
    .AddCheck("self", () => HealthCheckResult.Healthy());

// Configure CORS
var corsOrigins = builder.Configuration.GetSection("Cors:AllowedOrigins").Get<string[]>();
if (corsOrigins != null && corsOrigins.Length > 0)
{
    builder.Services.AddCors(options =>
    {
        options.AddDefaultPolicy(policy =>
            policy.WithOrigins(corsOrigins)
                  .AllowAnyHeader()
                  .AllowAnyMethod());
    });
}

// OpenTelemetry
var tracingOtlpEndpoint = builder.Configuration.GetValue<string>("Telemetry:OtlpEndpoint");
var otel = builder.Services.AddOpenTelemetry();

otel.ConfigureResource(resource => resource.AddService(serviceName: builder.Environment.ApplicationName));

var greeterMeter = new Meter("HelloApi.GreeterMeter", "1.0.0");
var greeterActivitySource = new ActivitySource("HelloApi.GreeterActivitySource");
var totalGreetingsCounter = greeterMeter.CreateCounter<int>("greetings.count", description: "Counts the number of greetings");
var greetingsByLangCounter = greeterMeter.CreateCounter<int>("greetings.by_lang", description: "Counts the number of greetings by language");

otel.WithMetrics(metrics =>
    metrics
        .AddAspNetCoreInstrumentation()
        .AddMeter(greeterMeter.Name)
        .AddMeter("Microsoft.AspNetCore.Hosting")
        .AddMeter("Microsoft.AspNetCore.Server.Kestrel")
        .AddMeter("System.Net.Http")
        .AddMeter("System.Net.NameResolution")
        .AddPrometheusExporter()
);

otel.WithTracing(tracing =>
{
    tracing.AddAspNetCoreInstrumentation();
    tracing.AddHttpClientInstrumentation();
    tracing.AddSource(greeterActivitySource.Name);
    if (!string.IsNullOrWhiteSpace(tracingOtlpEndpoint))
    {
        tracing.AddOtlpExporter(otlpOptions =>
        {
            otlpOptions.Endpoint = new Uri(tracingOtlpEndpoint);
        });
    }
    else
    {
        tracing.AddConsoleExporter();
    }
});


var app = builder.Build();

// Middleware
if (!app.Environment.IsDevelopment())
{
    app.UseHttpsRedirection();
}

if (app.Environment.IsDevelopment() || app.Environment.IsStaging())
{
    
    app.Logger.LogInformation("Loaded environment: {Env}, OTLP: {Endpoint}",
        app.Environment.EnvironmentName,
        app.Configuration["Tracing:OtlpEndpoint"]);

    app.MapOpenApi();
}

app.UseHealthChecks("/health");
app.MapHealthChecks("/health/live");
app.MapHealthChecks("/health/ready");

app.MapPrometheusScrapingEndpoint();

// Apply CORS only if configured
if (corsOrigins != null && corsOrigins.Length > 0)
{
    app.UseCors();
}


app.MapGet("/", (string? name, string? lang) =>
{
    using var activity = greeterActivitySource.StartActivity("GenerateGreeting");

    activity?.SetTag("request.query.name", name ?? "null");
    activity?.SetTag("request.query.lang", lang ?? "null");

    name ??= "World";
    lang ??= "en";

    var greeting = lang.ToLower() switch
    {
    "es" => "¡Hola",
    "fr" => "Bonjour",
     _ => "Hello"
    };
    var message = $"{greeting} {name}!";

    totalGreetingsCounter.Add(1);
    greetingsByLangCounter.Add(1, new KeyValuePair<string, object?>("lang", lang));

    activity?.AddEvent(new ActivityEvent("GreetingGenerated", tags: new ActivityTagsCollection
    {
        { "message", message }
    }));

    // Return structured JSON
    var response = new GreetingResponse(message, lang, DateTime.UtcNow);
    return Results.Ok(response);
}).WithName("GetHelloWorld");

// Lifecycle logs
app.Lifetime.ApplicationStarted.Register(() =>
{
    app.Logger.LogInformation("Application has started successfully.");
});

app.Lifetime.ApplicationStopping.Register(() =>
{
    app.Logger.LogInformation("Application is shutting down gracefully...");
});

app.Run();

/// Define record for structured response
/// <summary>
/// Represents a greeting response.
/// </summary>
record GreetingResponse(string Message, string Language, DateTime Timestamp);

/// Partial Program class for testing purposes
public partial class Program { }

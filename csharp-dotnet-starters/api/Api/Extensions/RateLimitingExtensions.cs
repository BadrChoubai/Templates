using System.Security.Claims;
using System.Threading.RateLimiting;

using Microsoft.AspNetCore.RateLimiting;
using Microsoft.Extensions.Options;

namespace Api.Extensions
{
    public static class RateLimitingExtensions
    {
        private const string Policy = "PerUserRatelimit";

        public static void AddRateLimiting(this IServiceCollection services)
        {
            services.AddRateLimiter();

            services.AddOptions<TokenBucketRateLimiterOptions>()
                .Configure(options =>
                {
                    // Defaults; override in the "RateLimiting" configuration section.
                    options.ReplenishmentPeriod = TimeSpan.FromSeconds(60);
                    options.AutoReplenishment = true;
                    options.TokenLimit = 100;
                    options.TokensPerPeriod = 100;
                    options.QueueLimit = 0; // reject with 429 instead of holding requests
                }).BindConfiguration("RateLimiting");

            services.AddOptions<RateLimiterOptions>()
                .Configure((RateLimiterOptions options,
                    IOptionsMonitor<TokenBucketRateLimiterOptions> perUserRateLimiterOptions) =>
                {
                    options.RejectionStatusCode = StatusCodes.Status429TooManyRequests;

                    options.AddPolicy(Policy, context =>
                    {
                        // Authenticated callers get their own bucket; anonymous callers share one per IP.
                        var partitionKey = context.User.FindFirstValue(ClaimTypes.NameIdentifier)
                            ?? context.Connection.RemoteIpAddress?.ToString()
                            ?? "anonymous";
                        return RateLimitPartition.GetTokenBucketLimiter(partitionKey,
                            _ => perUserRateLimiterOptions.CurrentValue);
                    });
                });
        }

        public static void RequirePerUserRateLimit(this IEndpointConventionBuilder builder)
        {
            builder.RequireRateLimiting(Policy);
        }
    }
}

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
                    options.ReplenishmentPeriod = TimeSpan.FromSeconds(180);
                    options.AutoReplenishment = true;
                    options.TokenLimit = 10;
                    options.TokensPerPeriod = 10;
                    options.QueueLimit = 100;
                }).BindConfiguration("RateLimiting");

            services.AddOptions<RateLimiterOptions>()
                .Configure((RateLimiterOptions options,
                    IOptionsMonitor<TokenBucketRateLimiterOptions> perUserRateLimiterOptions) =>
                {
                    options.RejectionStatusCode = StatusCodes.Status429TooManyRequests;

                    options.AddPolicy(Policy, context =>
                    {
                        var username = context.User.FindFirstValue(ClaimTypes.NameIdentifier);
                        Console.WriteLine(username);
                        return RateLimitPartition.GetTokenBucketLimiter(username,
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
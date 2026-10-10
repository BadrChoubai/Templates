using Api.Model;
using Api.Repository;

namespace Api
{
    public class DataSeeder(ApiDbContext context)
    {
        public void Seed()
        {
            if (context.Items.Any())
            {
                return;
            }

            var items = new List<Item>
            {
                new()
                {
                    Id = new Guid("c04ab9a0-e398-471d-8cfe-95317ccf0a20"),
                    Name = "First item",
                    Description = "An example item"
                },
                new()
                {
                    Id = new Guid("1e95757d-066d-41a6-b5d2-5e34b71d6051"),
                    Name = "Second item",
                    Description = "Another example item"
                }
            };

            context.Items.AddRange(items);
            context.SaveChanges();
        }
    }
}

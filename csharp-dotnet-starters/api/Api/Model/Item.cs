using System.ComponentModel.DataAnnotations;

namespace Api.Model;

public class Item
{
    public Guid Id { get; init; } = Guid.NewGuid();

    [MaxLength(255)] public string Name { get; set; } = string.Empty;

    [MaxLength(1024)] public string Description { get; set; } = string.Empty;
}

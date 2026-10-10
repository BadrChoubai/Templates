using System.ComponentModel.DataAnnotations;

namespace Api.Model;

/// <summary>
/// Request body for creating an item. Validated automatically by <c>AddValidation</c>.
/// </summary>
public record CreateItemRequest
{
    [Required(ErrorMessage = "name is required")]
    [MaxLength(255, ErrorMessage = "name must be at most 255 characters")]
    public string? Name { get; init; }

    [MaxLength(1024, ErrorMessage = "description must be at most 1024 characters")]
    public string? Description { get; init; }
}

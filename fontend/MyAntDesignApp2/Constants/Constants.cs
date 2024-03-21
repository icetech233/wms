using System.Text.Json;

namespace WmsApp
{
    public class Constants
    {
        // 
        public static JsonSerializerOptions CamelCaseOpt = new JsonSerializerOptions { PropertyNamingPolicy = JsonNamingPolicy.CamelCase };

    }
}

using System.Text.Json;

namespace MyAntDesignApp2
{
    public class Constants
    {
        // 
        public static JsonSerializerOptions CamelCaseOpt = new JsonSerializerOptions { PropertyNamingPolicy = JsonNamingPolicy.CamelCase };

    }
}

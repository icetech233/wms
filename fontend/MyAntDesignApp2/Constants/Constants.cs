using System.Text.Json;

namespace WmsApp
{
    public class Constants
    {
        // 
        public static JsonSerializerOptions CamelCaseOpt = new JsonSerializerOptions { PropertyNamingPolicy = JsonNamingPolicy.CamelCase };

        // mr-2	margin-right: 0.5rem; /* 8px */  margin-right: 8px;
        public const string mr2 = "margin-right: 0.5rem;";
    }
}

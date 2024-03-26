using System;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading.Tasks;
using WmsApp.Models;

namespace WmsApp.Services
{
    public interface IDictService
    {
        Task<DictResp> GetDictListAsync();
    }

    public class DictService : IDictService
    {
        private readonly HttpClient _httpClient;
        public DictService(HttpClient c)
        {
            _httpClient = c;
        }
        public async Task<DictResp> GetDictListAsync() {
            string requestUri = "/api/v1/dict/list?s=" + Random.Shared.Next(int.MaxValue);
            return await _httpClient.GetFromJsonAsync<DictResp>(requestUri);
        }

    }
}

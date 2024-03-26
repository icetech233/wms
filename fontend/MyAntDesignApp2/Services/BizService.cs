using System;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading.Tasks;
using WmsApp.Models;

namespace WmsApp.Services
{
    public interface IBizService
    {
        Task<DictResp> GetDictListAsync();
    }

    public class BizService : IBizService
    {
        private readonly HttpClient _httpClient;
        public BizService(HttpClient c)
        {
            _httpClient = c;
        }

        public async Task<DictResp> GetDictListAsync()
        {
            string requestUri = "/api/v1/dict/list?s=" + Random.Shared.Next(int.MaxValue);
            return await _httpClient.GetFromJsonAsync<DictResp>(requestUri);
        }
    }
}

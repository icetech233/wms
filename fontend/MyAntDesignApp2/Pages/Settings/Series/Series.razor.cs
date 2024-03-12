using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading.Tasks;
using AntDesign;
using Microsoft.AspNetCore.Components;
using MyAntDesignApp2.Models;

namespace MyAntDesignApp2.Pages.Settings
{

    public partial class Series
    {
        const string _PageKey = "网关类型:";

        [Inject]
        public HttpClient hc { get; set; }

        public Spu[] spuList;

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();
            await Refresh();
        }

        private async Task Refresh()
        {
            // /api/v1/spu/list
            string requestUri = "/api/v1/spu/list?s=" + Random.Shared.Next(int.MaxValue);
            SpuResp resp = await hc.GetFromJsonAsync<SpuResp>(requestUri);
            Console.WriteLine($"refresh code {resp.Code} count {resp.Data?.Length}");
            if (resp.Code != 200)
            {
                await Notice.Error(_notice, resp.Msg);
            }
            // spu
        }

    }
}

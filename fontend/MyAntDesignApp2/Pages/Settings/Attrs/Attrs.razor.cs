using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Runtime.Intrinsics.X86;
using System.Threading.Tasks;
using AntDesign;
using Microsoft.AspNetCore.Components;
using MyAntDesignApp2.Models;

namespace MyAntDesignApp2.Pages.Settings
{

    public partial class Attrs
    {
        private string _attrPageKey = "规格如下:";

        Attr[] attrList;

        [Inject]
        public HttpClient hc { get; init; }

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();

            string requestUri = "http://hw.acgzj.cn:8081" +
            "/api/v1/attr/list?s" + +Random.Shared.Next(int.MaxValue);
            var resp = await hc.GetFromJsonAsync<AttrResp>(requestUri);

            Console.WriteLine(resp);
            if (resp.Code != 200)
            {
                // 弹窗 resp.Msg;
                return;
            }

            attrList = resp.Data;

            requestUri = "http://hw.acgzj.cn:8081" +
"/api/v1/menu/list?s" + +Random.Shared.Next(int.MaxValue);
            var resp2 = await hc.GetFromJsonAsync<MenuResp>(requestUri);

            Console.WriteLine(resp2);

        }
    }
}

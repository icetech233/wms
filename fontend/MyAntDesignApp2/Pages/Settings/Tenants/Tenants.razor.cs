using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text.Json;
using System.Threading.Tasks;
using AntDesign;
using Microsoft.AspNetCore.Components;
using MyAntDesignApp2.Models;

namespace MyAntDesignApp2.Pages.Settings
{
    public partial class Tenants
    {
        const string _PageKey = "企业信息:";
        string size = "middle";
        bool _visible = false;
        string modalTitle = "添加"; // 编辑
        private Tenant[] tenantList;

        [Inject]
        public HttpClient hc { get; set; }

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();

            await Refresh();
        }

        private TenantModel model = new();
        public class TenantModel
        {
        }

        private async Task Refresh()
        {
            string requestUri = "/api/v1/tenant/list?s=" + Random.Shared.Next(int.MaxValue);
            TenantResp resp = await hc.GetFromJsonAsync<TenantResp> (requestUri);
            Console.WriteLine($"refresh code {resp.Code} count {resp.Data?.Length}");
            if (resp.Code != 200)
            {
                await Notice.Error(_notice, resp.Msg);
            }
            tenantList = resp.Data;
        }

        private void Delete(long id)
        {
            // 请求网络 delete 接口
            // warehouseList = warehouseList.Where(x => x.WarehouseID != id).ToArray();
        }

    }
}

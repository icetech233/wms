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
using WmsApp.Models;

namespace WmsApp.Pages.Settings
{
    public partial class Tenants
    {
        const string _PageKey = "企业信息:";
        string size = "middle";
        bool _visible = false;
        string modalTitle = "添加"; // 编辑

        private Tenant[] tenantAllList;
        private Tenant[] tenantList;

        [Inject]
        public HttpClient hc { get; set; }

        public EventCallback<string> OnChangeHandle;

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();
            await RefreshTableData();

            OnChangeHandle = new EventCallback<string>(null, () =>
            {
                RefreshTableDataWithEntType(bindEntType);
            });
        }

        private TenantModel model = new();
        public class TenantModel
        {
        }

        private async Task RefreshTableData()
        {
            string requestUri = $"/api/v1/tenant/list?ent_type={bindEntType}&s={Random.Shared.Next(int.MaxValue)}";
            TenantResp resp = await hc.GetFromJsonAsync<TenantResp>(requestUri);
            Console.WriteLine($"refresh code {resp.Code} count {resp.Data?.Length}");
            if (resp.Code != 200)
            {
                await Notice.Error(_notice, resp.Msg);
            }
            if (tenantAllList == null || tenantAllList.Length == 0)
            {
                tenantAllList = resp.Data;
            }
            else if ("all" == bindEntType) 
            {
                tenantAllList = resp.Data;
                tenantList = [.. tenantAllList];
            }
            // 分割
            if ("supplier" == bindEntType)
            {
                tenantList = tenantAllList.Where(a => a.IsSupplier).ToArray();
            }
            else if ("customer" == bindEntType)
            {
                tenantList = tenantAllList.Where(a => a.IsCustomer).ToArray();
            }

        }

        private void RefreshTableDataWithEntType(string ent_type)
        {
            if ("supplier" == ent_type)
            {
                tenantList = tenantAllList.Where(a => a.IsSupplier).ToArray();
            }
            else if ("customer" == ent_type)
            {
                tenantList = tenantAllList.Where(a => a.IsCustomer).ToArray();
            }
            else {
                tenantList = [.. tenantAllList];
            }
        }


        private void Delete(long id)
        {
            // 请求网络 delete 接口
            // warehouseList = warehouseList.Where(x => x.WarehouseID != id).ToArray();
        }

    }
}

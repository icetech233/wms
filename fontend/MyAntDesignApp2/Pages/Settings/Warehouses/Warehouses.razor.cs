using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text.Json;
using System.Threading.Tasks;
using AntDesign;
using Microsoft.AspNetCore.Components;
using Microsoft.Extensions.Logging;
using MyAntDesignApp2.Models;

namespace MyAntDesignApp2.Pages.Settings
{
    public partial class Warehouses
    {
        const string _PageKey = "规格如下:";
        string size = "middle";
        bool _visible = false;
        string modalTitle = "添加仓库"; // 编辑仓库

        [Inject]
        private ILogger<Warehouses> logger { get; set; }
        [Inject]
        public HttpClient hc { get; init; }

        protected override async Task OnInitializedAsync()
        {
            logger.LogWarning("OnInitializedAsync has clicked me!");
            await base.OnInitializedAsync();

            await Refresh();
        }

        private WarehouseModel model = new();

        public class WarehouseModel
        {
            public long WarehouseID { get; set; }

            /// <summary>
            /// 1 启用 0 禁用
            /// </summary>
            public bool Enable { get; set; }
            public string WarehouseCode { get; set; }
            public string WarehouseName { get; set; }

            public string Supervisor { get; set; }
            public string Tel { get; set; }
            [Required]
            public string Address { get; set; }
        }

        private void AddWarehouses()
        {
            modalTitle = "添加仓库";
            _visible = true;

            model.WarehouseID = -1;
            model.Enable = false;
            model.WarehouseCode = "";
            model.WarehouseName = "";
            model.Supervisor = "";
            model.Tel = "";
            model.Address = "";

        }

        private async void EditWarehouse(Warehouse w)
        {
            modalTitle = "编辑仓库";
            if (w == null)
            {
                await Notice.Error(_notice, "无数据");
                return;
            }

            model.WarehouseID = w.WarehouseID;
            model.Enable = w.Enable;
            model.WarehouseCode = w.WarehouseCode;
            model.WarehouseName = w.WarehouseName;
            model.Supervisor = w.Supervisor;
            model.Tel = w.Tel;
            model.Address = w.Address;

            _visible = true;
        }


        private async Task Refresh()
        {
            string requestUri = "http://hw.acgzj.cn:8081" +
            "/api/v1/warehouse/list?s=" + Random.Shared.Next(int.MaxValue);
            WarehouseResp resp = await hc.GetFromJsonAsync<WarehouseResp>(requestUri);
            if (resp.Code != 200)
            {
                await Notice.Error(_notice, resp.Msg);
                return;
            }
            warehouseList = resp.Data;
        }

        private void Delete(long id)
        {
            // 请求网络 delete 接口
            warehouseList = warehouseList.Where(x => x.WarehouseID != id).ToArray();
        }

    }
}

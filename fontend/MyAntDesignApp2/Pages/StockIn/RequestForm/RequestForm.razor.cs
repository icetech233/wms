﻿using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text.Json;
using System.Threading.Tasks;
using AntDesign;
using Blazored.LocalStorage;
using Microsoft.AspNetCore.Components;
using WmsApp.Models;
using WmsApp.Services;

namespace WmsApp.Pages.StockIn
{
    public partial class RequestForm
    {
        private string requestFormTitle = "申请入库";
        // private string requestFormTitle = "申请入库";
        private StockAddModel stockModel = new();

        [Inject]
        public IBizService BizSrv { get; set; }
        [Inject]
        public ILocalStorageService LocalStorage { get; set; }

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();

            DictResp dictResp = await BizSrv.GetDictListAsync();
            WarehouseResp warehouseResp = await BizSrv.GetWarehouseListAsync();

            // 入库类型 = 1
            _stockInTypes = dictResp.Data.Where(a => a.BizType == 1).ToList();
            
            _warehouses = warehouseResp.Data.ToList();


            await Console.Out.WriteAsync("OnInitializedAsync");
        }

    }
}

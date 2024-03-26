using AntDesign.ProLayout;
using System;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Microsoft.Extensions.DependencyInjection;
using AntDesign;
using WmsApp.Services;
using Blazored.LocalStorage;

namespace WmsApp
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var builder = WebAssemblyHostBuilder.CreateDefault(args);
            builder.RootComponents.Add<App>("#app");

            // /data/menu.json
            // BaseAddress = new Uri(builder.HostEnvironment.BaseAddress)
            builder.Services.AddScoped(sp => new HttpClient
            {
                BaseAddress = new Uri("http://hw.acgzj.cn:8081/")
            });
            //
            builder.Services.AddBlazoredLocalStorage();
            builder.Services.AddAntDesign();
            builder.Services.Configure<ProSettings>(builder.Configuration.GetSection("ProSettings"));
            //builder.Services.AddScoped<IUserService, UserService>();
            builder.Services.AddScoped<IBizService, BizService>();
            builder.Services.AddScoped<IDictService, DictService>();
            await builder.Build().RunAsync();
        }
    }
}
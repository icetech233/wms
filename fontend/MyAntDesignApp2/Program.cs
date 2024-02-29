using AntDesign.ProLayout;
using System;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Microsoft.Extensions.DependencyInjection;

namespace MyAntDesignApp2
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var builder = WebAssemblyHostBuilder.CreateDefault(args);
            builder.RootComponents.Add<App>("#app");
            string baddr = builder.HostEnvironment.BaseAddress;

            builder.Services.AddScoped(sp => new HttpClient {
                BaseAddress = new Uri(baddr)
            });
            // HttpClient : HttpMessageInvoker
            // builder.Services.AddSingleton<HttpMessageInvoker, HttpClient>();
            builder.Services.AddAntDesign();
            builder.Services.Configure<ProSettings>(builder.Configuration.GetSection("ProSettings"));

            await builder.Build().RunAsync();
        }
    }
}
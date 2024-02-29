using System.Collections.Generic;
using System.Threading.Tasks;
using AntDesign.ProLayout; 
using Microsoft.AspNetCore.Components;

namespace MyAntDesignApp2.Pages.Profile {
  public partial class Advanced {


        private string _selectKey = "base";
        private readonly IList<TabPaneItem> _tabList = new List<TabPaneItem>
        {
            new TabPaneItem {Key = "detail", Tab = "Details"},
            new TabPaneItem {Key = "rules", Tab = "Rules"}
        };

    //private AdvancedProfileData _data = new AdvancedProfileData();

    //[Inject] protected IProfileService ProfileService { get; set; }

    protected override async Task OnInitializedAsync() {
      await base.OnInitializedAsync();
      //_data = await ProfileService.GetAdvancedAsync();
    }
  }
}
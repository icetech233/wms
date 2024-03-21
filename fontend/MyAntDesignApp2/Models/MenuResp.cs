namespace WmsApp.Models
{
    public class MenuResp : BaseResp
    {
        public object Debug { get; set; }
        public Menu[] Data { get; set; }
    }

}

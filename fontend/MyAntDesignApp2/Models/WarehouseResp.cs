namespace WmsApp.Models
{
    public class WarehouseResp : BaseResp
    {
        public object Debug { get; set; }
        public Warehouse[] Data { get; set; }
    }
}

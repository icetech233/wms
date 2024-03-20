namespace MyAntDesignApp2.Models
{
    public class WarehouseResp : BaseResp
    {
        public object Debug { get; set; }
        public Warehouse[] Data { get; set; }

    }

    public class DictResp : BaseResp
    {
        public object Debug { get; set; }
        public Dict[] Data { get; set; }
    }

}

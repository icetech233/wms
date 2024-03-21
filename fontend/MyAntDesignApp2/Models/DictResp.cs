namespace WmsApp.Models
{
    public class DictResp : BaseResp
    {
        public object Debug { get; set; }
        public Dict[] Data { get; set; }
    }

    public class TenantResp : BaseResp
    {
        public object Debug { get; set; }
        public Tenant[] Data { get; set; }
    }
}

using System.ComponentModel;

namespace WmsApp.Models
{
    public class AttrResp : BaseResp
    {
        public object Debug { get; set; }
        public Attr[] Data { get; set; }
    }
}

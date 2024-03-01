namespace MyAntDesignApp2.Models
{
    public class AttrResp : BaseResp
    {
        public object Debug { get; set; }
        public Attr[] Data { get; set; }
    }

    public class BaseResp
    {
        public int Code { get; init; }
        public string Msg { get; init; }
    }

}

namespace MyAntDesignApp2.Models
{
    public record class BaseResp(int Code, string Msg, object Debug, object Data);
    public record class AttrResp(int Code, string Msg, object Debug, Attr[] Data);
}

﻿using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class AttrResp : BaseResp
    {
        public object Debug { get; set; }
        public Attr[] Data { get; set; }
    }

    public class BaseResp
    {
        public int Code { get; set; }
        public string Msg { get; set; }
    }
}

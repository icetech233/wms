﻿using System.Collections.Generic;

namespace WmsApp.Models
{
    public class SpuResp : BaseResp
    {
        public object Debug { get; set; }
        public List<Spu> Data { get; set; }
    }
}

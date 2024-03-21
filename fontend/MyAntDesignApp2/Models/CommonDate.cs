using System.ComponentModel;

namespace WmsApp.Models
{
    public class CommonDate {
        [DisplayName("创建时间")]
        public string CreatedAt { get; set; }
        [DisplayName("修改时间")]
        public string UpdatedAt { get; set; }
    }
}

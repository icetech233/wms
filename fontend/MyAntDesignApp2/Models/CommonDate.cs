using System.ComponentModel;

namespace MyAntDesignApp2.Models
{
    public class CommonDate {
        [DisplayName("创建")]
        public string CreatedAt { get; set; }
        [DisplayName("修改")]
        public string UpdatedAt { get; set; }
    }
}

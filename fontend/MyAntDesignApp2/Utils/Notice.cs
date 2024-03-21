using AntDesign;
using System.Threading.Tasks;

namespace WmsApp
{
    public class Notice
    {
        public static async Task Error(NotificationService _notice, string msg)
        {
            await _notice.Open(new NotificationConfig()
            {
                Message = "错误",
                Description = msg,
                NotificationType = NotificationType.Error
            });
        }

    }
}

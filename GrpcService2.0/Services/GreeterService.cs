using System;
using System.Collections.Generic;
using System.Data;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Core;
using GrpcService1;
using Microsoft.Extensions.Logging;

namespace GrpcService2._0 {
    public class GreeterService : Greeter.GreeterBase, IObserver {

        public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context) {
            return Task.FromResult(new HelloReply {
                Message = "Hello " + request.Name
            });
        }

        public IServerStreamWriter<NumReply> responseStream;
        public ServerCallContext context;

        Random rand = new Random();

        public delegate void CheckNumber();
        public event CheckNumber Upd;

        static Number num = null;
        bool b;

        public override async Task GetRandNum(NumRequest request, IServerStreamWriter<NumReply> responseStream, ServerCallContext context) {

            this.responseStream = responseStream;
            this.context = context;
            if (num == null) {
                num = new Number();  
            }
            num.RegisterObserver(this);
            Console.ForegroundColor = ConsoleColor.Green;
            Console.WriteLine(" ÎËÂÌÚ ÔË¯∏Î");
            Console.ForegroundColor = ConsoleColor.Gray;
            Console.WriteLine(num);

            while (true) {
                //if (b) {
                //    var resp = new NumReply {
                //        Message = Convert.ToString(num.ValueNumber)
                //    };
                //    Console.WriteLine(resp + " ›“Œ »« ÀŒ√¿");
                //    await responseStream.WriteAsync(resp);
                //    b = false;
                //}

            }
            num.RemoveObserver(this);
        }

        public void Update(object ob) {
            Notification value = (Notification)ob;
            var resp = new NumReply {
                Message = Convert.ToString(value.number.ValueNumber)
            };
            try {
                responseStream.WriteAsync(resp);
            } catch (System.InvalidOperationException e) {
                Console.ForegroundColor = ConsoleColor.Red;
                Console.WriteLine(" ÎËÂÌÚ Û¯∏Î");
                Console.ForegroundColor = ConsoleColor.Gray;
                num.RemoveObserver(this);
            }
            
            b = true;
            //Upd?.Invoke(value);
            Console.WriteLine("”‡‡‡‡‡‡‡‡‡‡‡‡‡‡‡‡‡, Ò‡·ÓÚ‡Î ¿Ô‰ÂÈÚ");
        }
    }
}

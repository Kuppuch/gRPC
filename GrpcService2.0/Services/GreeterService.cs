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

        //public IServerStreamWriter<NumReply> responseStream;
        //public ServerCallContext context;

        string number;

        public delegate void CheckNumber();
        public event CheckNumber Upd;

        static Number num = null;
        bool b;

        public override async Task GetRandNum(NumRequest request, IServerStreamWriter<NumReply> responseStream, ServerCallContext context) {

            if (num == null) {
                num = new Number();  
            }
            num.RegisterObserver(this);
            Console.ForegroundColor = ConsoleColor.Green;
            Console.WriteLine(" ÎËÂÌÚ ÔË¯∏Î");
            Console.ForegroundColor = ConsoleColor.Gray;
            Console.WriteLine(num);

            while (true) {
                if (b) {
                    var resp = new NumReply {
                        Message = Convert.ToString(number)
                    };
                    Console.WriteLine(resp + " ›“Œ »« ÀŒ√¿");
                    await responseStream.WriteAsync(resp);
                    b = false;
                }

            }
        }

        public void Update(object ob) {
            Notification value = (Notification)ob;
            number = value.number.ValueNumber + " - " + value.i;
            b = true;
        }
    }
}

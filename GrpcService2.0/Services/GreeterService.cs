using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Core;
using GrpcService1;
using Microsoft.Extensions.Logging;

namespace GrpcService2._0 {
    public class GreeterService : Greeter.GreeterBase {
        /*private readonly ILogger<GreeterService> _logger;
        public GreeterService(ILogger<GreeterService> logger) {
            _logger = logger;
        }*/

        public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context) {
            return Task.FromResult(new HelloReply {
                Message = "Hello " + request.Name
            });
        }

        Random rand = new Random();

        public override async Task GetRandNum(NumRequest request, IServerStreamWriter<NumReply> responseStream, ServerCallContext context) {

            //int val = Convert.ToInt32(request.Name);

            while (true) {
                var resp = new NumReply {
                    Message = Convert.ToString(Program.v)
                };
                Console.WriteLine(resp + " ›“Œ »« ÀŒ√¿");
                await responseStream.WriteAsync(resp);
                Thread.Sleep(2000);
            }
        }
    }
}

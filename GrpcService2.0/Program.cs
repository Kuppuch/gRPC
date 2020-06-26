using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Hosting;

namespace GrpcService2._0 {
    public class Program {
        static Thread common = new Thread(new ThreadStart(Common));
        static Random rand = new Random();
        public static int v = 0;
        public static void Main(string[] args) {

            common.Name = "Общий поток";
            common.Start();

            CreateHostBuilder(args).Build().Start();
            
            
            
        }

        // Additional configuration is required to successfully run gRPC on macOS.
        // For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682
        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder => {
                    webBuilder.UseStartup<Startup>();
                });

        



        public static void Common() {
            while (true) {
                v = rand.Next(0, 20);
                //Console.WriteLine("___ " + v);
                Thread.Sleep(500 + rand.Next(0, 500));
            }
        }
    }
}

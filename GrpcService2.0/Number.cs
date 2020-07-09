using Grpc.Core;
using GrpcService1;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

namespace GrpcService2._0{
    public class Number : IObservable {

        private readonly object locker = new object();
        List<IObserver> observers = new List<IObserver>();
        List<IObserver> observers2Remove = new List<IObserver>();
        public int ValueNumber { get; set; }
        int num;

        public Number() {
            Thread common = new Thread(new ThreadStart(Generate)) {
                Name = "Общий поток"
            };
            common.Start();
        }

        public void RegisterObserver(IObserver o) {
            lock (locker) {
                observers.Add(o);
            }
        }

        public void RemoveObserver(IObserver o) {
            observers2Remove.Add(o);
        }

        public void NotifyObservers() {
            lock (locker) {
                foreach (IObserver o in observers2Remove) {
                    observers.Remove(o);
                }
                observers2Remove.Clear();
                int j = 0;
                foreach (IObserver o in observers) {
                    o.Update(new Notification() { observer = o, number = this, i = Convert.ToString(j) });
                    j++;
                }
            }
        }

        public void Generate() {
             
                Random rand = new Random();
                while (true) {
                    try {
                        num = rand.Next(0, 20);
                        ValueNumber = num;
                        //Console.WriteLine("Значение числа =  " + ValueNumber);
                        NotifyObservers();
                        Thread.Sleep(100 + rand.Next(0, 50));
                    } catch (Exception e) {
                        Console.WriteLine("              " + e);
                    }
                }
            
        }
    }

}

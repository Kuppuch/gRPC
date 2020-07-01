using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace GrpcService2._0 {

    public interface IObserver {
        void Update(Object ob);
    }
}

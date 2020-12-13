using EquationSolver.Domain;
using System.Threading.Tasks;

namespace EquationSolver
{
    public interface IEquationsDatabase
    {
        Task<EquationRoots> GetIfExist(Coefficients coefficients);
        Task<bool> Save(EquationRoots roots);
    }
}
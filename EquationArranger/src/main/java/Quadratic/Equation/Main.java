package Quadratic.Equation;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main
{
    public static void main(String[] args)
    {
        ApplicationContext context = new ClassPathXmlApplicationContext("Beans.xml");

        EquationArranger equationArranger = (EquationArranger) context.getBean("equationArranger");
        ArrangeResult arrangeResult = equationArranger.arrange(args[0]);

        if (arrangeResult.isArrangeSucceeded())
            System.out.println(arrangeResult.arrangedEquation());
        else
            System.out.println(arrangeResult.getErrorsAndWarnings());
    }
}
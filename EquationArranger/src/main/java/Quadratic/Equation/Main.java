package Quadratic.Equation;

import org.apache.log4j.PropertyConfigurator;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main
{
    public static void main(String[] args)
    {
        PropertyConfigurator.configure("log4j.properties");

        ApplicationContext context = new ClassPathXmlApplicationContext("Beans.xml");

        EquationArranger equationArranger = (EquationArranger) context.getBean("equationArranger");
        ArrangeResult arrangeResult = equationArranger.arrange("3x^2 + 6x -5 = 0");

        if (arrangeResult.isArrangeSucceeded())
            System.out.println(arrangeResult.arrangedEquation());
        else
            System.out.println(arrangeResult.getErrorsAndWarnings());
    }
}
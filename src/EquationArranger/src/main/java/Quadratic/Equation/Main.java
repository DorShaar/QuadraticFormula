package Quadratic.Equation;

import Quadratic.Equation.Arranger.ArrangeResult;
import Quadratic.Equation.Arranger.EquationArranger;
import Quadratic.Equation.Arranger.IEquationArranger;
import org.apache.log4j.Logger;
import org.apache.log4j.PropertyConfigurator;
import org.json.simple.JSONObject;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main
{
    final static Logger logger = Logger.getLogger(Main.class);

    public static void main(String[] args)
    {
        PropertyConfigurator.configure("log4j.properties");

        ApplicationContext context = new ClassPathXmlApplicationContext("Beans.xml");

        IEquationArranger equationArranger = (EquationArranger) context.getBean("equationArranger");

        long startTime = System.currentTimeMillis();

        String equation = "3x^2 + 6x -5 = 6x ^ 2 - 4x + 6";

        ArrangeResult arrangeResult = equationArranger.arrange(equation);

        long endArrangmentTime = System.currentTimeMillis();
        logger.info("Arrangement for equation " + equation + " took " + (endArrangmentTime - startTime) + "ms");

        if (arrangeResult.isArrangeSucceeded())
        {
            logger.info(arrangeResult.arrangedEquation());
            sendMessage(arrangeResult, startTime, endArrangmentTime);
        }
        else
        {
            logger.info(arrangeResult.getErrorsAndWarnings());
        }

        long endSendMessageTime = System.currentTimeMillis();
        logger.info("Arrangement and send message for equation " + equation +
                " took " + (endSendMessageTime - startTime) + "ms");
    }

    private static void sendMessage(ArrangeResult arrangeResult, long startTime, long endTime)
    {
        JSONObject json = new JSONObject();
        json.put("originalEquation", arrangeResult.originalEquation());
        json.put("arrangedEquation", arrangeResult.arrangedEquation());
        json.put("startTime", startTime);
        json.put("endTime", endTime);

        logger.debug(json.toJSONString());
    }
}
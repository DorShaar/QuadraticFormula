package Quadratic.Equation;

import java.util.List;
import java.util.Vector;

public class ArrangeResult
{
    private String originalEquation;
    private String arrangedEquation;
    private Boolean isArrangeSucceeded;
    private List<String> notes;

    private ArrangeResult(String originalEquation, String arrangedEquation)
    {
        this.originalEquation = originalEquation;
        this.arrangedEquation = arrangedEquation;
        isArrangeSucceeded = true;
        notes = new Vector<>(0);
    }

    private ArrangeResult(String originalEquation,
                          String arrangedEquation,
                          List<String> notes)
    {
        this.originalEquation = originalEquation;
        this.arrangedEquation = arrangedEquation;
        isArrangeSucceeded = true;
        this.notes = notes;
    }

    private ArrangeResult(String originalEquation, List<String> errors)
    {
        this.originalEquation = originalEquation;
        this.arrangedEquation = null;
        isArrangeSucceeded = false;
        this.notes = errors;
    }

    public static ArrangeResult SuccessArrange(String originalEquation,
                                               String arrangedEquation)
    {
        return new ArrangeResult(originalEquation, arrangedEquation);
    }

    public static ArrangeResult SuccessArrangeWithNotes(String originalEquation,
                                                        String arrangedEquation,
                                                        List<String> notes)
    {
        return new ArrangeResult(originalEquation, arrangedEquation, notes);
    }

    public static ArrangeResult FailedArrange(String originalEquation,
                                                        List<String> notes)
    {
        return new ArrangeResult(originalEquation, notes);
    }

    public String arrangedEquation()
    {
        return arrangedEquation;
    }

    public Boolean isArrangeSucceeded()
    {
        return isArrangeSucceeded;
    }
}